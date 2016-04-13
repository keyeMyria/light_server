/**
 * Autogenerated by Thrift Compiler (0.9.2)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#ifndef proc_TYPES_H
#define proc_TYPES_H

#include <iosfwd>

#include <thrift/Thrift.h>
#include <thrift/TApplicationException.h>
#include <thrift/protocol/TProtocol.h>
#include <thrift/transport/TTransport.h>

#include <thrift/cxxfunctional.h>


namespace apache { namespace thrift { namespace test {

class MyError;

typedef struct _MyError__isset {
  _MyError__isset() : message(false) {}
  bool message :1;
} _MyError__isset;

class MyError : public ::apache::thrift::TException {
 public:

  static const char* ascii_fingerprint; // = "EFB929595D312AC8F305D5A794CFEDA1";
  static const uint8_t binary_fingerprint[16]; // = {0xEF,0xB9,0x29,0x59,0x5D,0x31,0x2A,0xC8,0xF3,0x05,0xD5,0xA7,0x94,0xCF,0xED,0xA1};

  MyError(const MyError&);
  MyError& operator=(const MyError&);
  MyError() : message() {
  }

  virtual ~MyError() throw();
  std::string message;

  _MyError__isset __isset;

  void __set_message(const std::string& val);

  bool operator == (const MyError & rhs) const
  {
    if (!(message == rhs.message))
      return false;
    return true;
  }
  bool operator != (const MyError &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const MyError & ) const;

  template <class Protocol_>
  uint32_t read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t write(Protocol_* oprot) const;

  friend std::ostream& operator<<(std::ostream& out, const MyError& obj);
};

void swap(MyError &a, MyError &b);

}}} // namespace

#include "proc_types.tcc"

#endif
