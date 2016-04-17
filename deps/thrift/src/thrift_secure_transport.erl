%%
%% Licensed to the Apache Software Foundation (ASF) under one
%% or more contributor license agreements. See the NOTICE file
%% distributed with this work for additional information
%% regarding copyright ownership. The ASF licenses this file
%% to you under the Apache License, Version 2.0 (the
%% "License"); you may not use this file except in compliance
%% with the License. You may obtain a copy of the License at
%%
%%   http://www.apache.org/licenses/LICENSE-2.0
%%
%% Unless required by applicable law or agreed to in writing,
%% software distributed under the License is distributed on an
%% "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
%% KIND, either express or implied. See the License for the
%% specific language governing permissions and limitations
%% under the License.
%%

-module(thrift_secure_transport).

-behaviour(thrift_transport).

%% API
-export([new/3]).

%% thrift_transport callbacks
-export([write/2, read/2, flush/1, close/1]).

-record(crypto_transport, {wrapped, % a thrift_transport
    aes_key,
    aes_vector,
	read_buffer, % iolist()
	write_buffer % iolist()
 }).
-type state() :: #crypto_transport{}.
-include("thrift_transport_behaviour.hrl").

new(WrappedTransport, AesKey, AesVector) ->
	State = #crypto_transport{wrapped = WrappedTransport,
                              aes_key = AesKey,
                              aes_vector = AesVector,
							  read_buffer = [],
							  write_buffer = []},
	thrift_transport:new(?MODULE, State).


%% Writes data into the buffer
write(State = #crypto_transport{write_buffer = WBuf}, Data) ->
	{State#crypto_transport{write_buffer = [WBuf, Data]}, ok}.

%% Flushes the buffer through to the wrapped transport
flush(State0 = #crypto_transport{write_buffer = Buffer,
	wrapped = Wrapped0}) ->

	Encrypted = encrypt(State0, Buffer),

	FrameLen = byte_size(Encrypted),
	Data     = [<<FrameLen:32/integer-signed-big>>, Encrypted],

	{Wrapped1, Response} = thrift_transport:write(Wrapped0, Data),

	{Wrapped2, _} = thrift_transport:flush(Wrapped1),

	State1 = State0#crypto_transport{wrapped = Wrapped2, write_buffer = []},
	{State1, Response}.

%% Closes the transport and the wrapped transport
close(State = #crypto_transport{wrapped = Wrapped0}) ->
	{Wrapped1, Result} = thrift_transport:close(Wrapped0),
	NewState = State#crypto_transport{wrapped = Wrapped1},
	{NewState, Result}.

%% Reads data through from the wrapped transport
read(State0 = #crypto_transport{wrapped = Wrapped0, read_buffer = RBuf},
	Len) when is_integer(Len) ->
	{Wrapped1, {RBuf1, RBuf1Size}} =
		%% if the read buffer is empty, read another frame
	%% otherwise, just read from what's left in the buffer
	case iolist_size(RBuf) of
		0 ->
			%% read the frame length
			case thrift_transport:read(Wrapped0, 4) of
				{WrappedS1, {ok, <<FrameLen:32/integer-signed-big, _/binary>>}} ->
					%% then read the data
					case thrift_transport:read(WrappedS1, FrameLen) of
						{WrappedS2, {ok, Bin}} ->
							Bin1 = decrypt(State0, Bin),
							{WrappedS2, {Bin1, erlang:byte_size(Bin1)}};
						{WrappedS2, {error, Reason1}} ->
							{WrappedS2, {error, Reason1}}
					end;
				{WrappedS1, {error, Reason2}} ->
					{WrappedS1, {error, Reason2}}
			end;
		Sz ->
			{Wrapped0, {RBuf, Sz}}
	end,

	%% pull off Give bytes, return them to the user, leave the rest in the buffer
	case RBuf1 of
		error ->
			{ State0#crypto_transport {wrapped = Wrapped1, read_buffer = [] },
				{RBuf1, RBuf1Size} };
		_ ->
			Give = min(RBuf1Size, Len),
			<<Data:Give/binary, RBuf2/binary>> = iolist_to_binary(RBuf1),

			{ State0#crypto_transport{wrapped = Wrapped1, read_buffer=RBuf2},
				{ok, Data} }
	end.

encrypt(#crypto_transport{aes_key=Key, aes_vector = Vector}, IOList) ->
%%     Key = get(aes_key),
%%     Vector = get(aes_vector),
    aes:encrypt(IOList, Key, Vector).

decrypt(#crypto_transport{aes_key=Key, aes_vector = Vector},Bytes) ->
%%     Key = get(aes_key),
%%     Vector = get(aes_vector),
    aes:decrypt(Bytes, Key, Vector).
