// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"stress"
	"strings"
	"thrift"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  void echoVoid()")
	fmt.Fprintln(os.Stderr, "  byte echoByte(byte arg)")
	fmt.Fprintln(os.Stderr, "  i32 echoI32(i32 arg)")
	fmt.Fprintln(os.Stderr, "  i64 echoI64(i64 arg)")
	fmt.Fprintln(os.Stderr, "  string echoString(string arg)")
	fmt.Fprintln(os.Stderr, "   echoList( arg)")
	fmt.Fprintln(os.Stderr, "   echoSet( arg)")
	fmt.Fprintln(os.Stderr, "   echoMap( arg)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := stress.NewServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "echoVoid":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "EchoVoid requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.EchoVoid())
		fmt.Print("\n")
		break
	case "echoByte":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "EchoByte requires 1 args")
			flag.Usage()
		}
		tmp0, err26 := (strconv.Atoi(flag.Arg(1)))
		if err26 != nil {
			Usage()
			return
		}
		argvalue0 := byte(tmp0)
		value0 := argvalue0
		fmt.Print(client.EchoByte(value0))
		fmt.Print("\n")
		break
	case "echoI32":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "EchoI32 requires 1 args")
			flag.Usage()
		}
		tmp0, err27 := (strconv.Atoi(flag.Arg(1)))
		if err27 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.EchoI32(value0))
		fmt.Print("\n")
		break
	case "echoI64":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "EchoI64 requires 1 args")
			flag.Usage()
		}
		argvalue0, err28 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err28 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.EchoI64(value0))
		fmt.Print("\n")
		break
	case "echoString":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "EchoString requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.EchoString(value0))
		fmt.Print("\n")
		break
	case "echoList":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "EchoList requires 1 args")
			flag.Usage()
		}
		arg30 := flag.Arg(1)
		mbTrans31 := thrift.NewTMemoryBufferLen(len(arg30))
		defer mbTrans31.Close()
		_, err32 := mbTrans31.WriteString(arg30)
		if err32 != nil {
			Usage()
			return
		}
		factory33 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt34 := factory33.GetProtocol(mbTrans31)
		containerStruct0 := stress.NewEchoListArgs()
		err35 := containerStruct0.ReadField1(jsProt34)
		if err35 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Arg
		value0 := argvalue0
		fmt.Print(client.EchoList(value0))
		fmt.Print("\n")
		break
	case "echoSet":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "EchoSet requires 1 args")
			flag.Usage()
		}
		arg36 := flag.Arg(1)
		mbTrans37 := thrift.NewTMemoryBufferLen(len(arg36))
		defer mbTrans37.Close()
		_, err38 := mbTrans37.WriteString(arg36)
		if err38 != nil {
			Usage()
			return
		}
		factory39 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt40 := factory39.GetProtocol(mbTrans37)
		containerStruct0 := stress.NewEchoSetArgs()
		err41 := containerStruct0.ReadField1(jsProt40)
		if err41 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Arg
		value0 := argvalue0
		fmt.Print(client.EchoSet(value0))
		fmt.Print("\n")
		break
	case "echoMap":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "EchoMap requires 1 args")
			flag.Usage()
		}
		arg42 := flag.Arg(1)
		mbTrans43 := thrift.NewTMemoryBufferLen(len(arg42))
		defer mbTrans43.Close()
		_, err44 := mbTrans43.WriteString(arg42)
		if err44 != nil {
			Usage()
			return
		}
		factory45 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt46 := factory45.GetProtocol(mbTrans43)
		containerStruct0 := stress.NewEchoMapArgs()
		err47 := containerStruct0.ReadField1(jsProt46)
		if err47 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Arg
		value0 := argvalue0
		fmt.Print(client.EchoMap(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
