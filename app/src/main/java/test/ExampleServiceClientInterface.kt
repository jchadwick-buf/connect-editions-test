// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: test/test.proto
//
package test

import com.connectrpc.Headers
import com.connectrpc.ResponseMessage

public interface ExampleServiceClientInterface {
  public suspend fun exampleCall(request: Test.ExampleMessage, headers: Headers = emptyMap()):
      ResponseMessage<Test.ExampleMessage>
}
