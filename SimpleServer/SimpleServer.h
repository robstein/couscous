// #include "Arduino.h"
// #include <ESP8266WiFi.h>

// // #define HTTP_STATUS_OK                                                         \
// //   String("HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nConnection: close\r\n\r\n")
// // #define HTTP_STATUS_INTERNAL_ERROR                                             \
// //   String("HTTP/1.1 500 OK\r\nContent-Type: text/html\r\nConnection: close\r\n\r\n")
// // #define HTTP_STATUS_NOT_FOUND                                                  \
// //   String("HTTP/1.0 404 Not Found\r\nContent-Type: text/html\r\nConnection: close\r\n\r\n")

// typedef struct {
//   int status;
//   String body;
// } response;

// typedef response *(*handleFn)(String, String);

// typedef struct {
//   String method;
//   String path;
//   handleFn handle;
// } handler;

// extern void SimpleServe(WiFiServer server, handler handlers[],
//                         int handlers_size);