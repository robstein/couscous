// #include "Arduino.h"
// #include <ESP8266WiFi.h>
// #include <SimpleServer.h>

// bool isFinalLine(String line) {
//   return (line.length() == 1 && line[0] == '\n');
// }
// bool isBlankLine(String line) {
//   return (line.length() == 2 && line.equals("\n\r"));
// }

// void SimpleServe(WiFiServer server, handler handlers[], int handlers_size) {
//   WiFiClient client = server.available();
//   if (client) {
//     handleFn fn = NULL;
//     String headers = "";
//     String body = "";

//     bool scanningfirstLine = true;
//     bool scanningHeaders = false;
//     bool scanningBody = false;

//     while (client.connected()) {
//       if (client.available()) {
//         String line = client.readStringUntil('\r');
//         Serial.print(line);

//         if (scanningfirstLine) {
//           for (int i = 0; i < handlers_size; i++) {
//             if (line.startsWith(handlers[i].method + " " + handlers[i].path))
//             {
//               fn = handlers[i].handle;
//               break;
//             }
//           }
//           scanningfirstLine = false;
//           scanningHeaders = true;
//           continue;
//         }

//         if (scanningHeaders) {
//           headers += line;
//           if ((line.length() == 2 && line.equals("\n\r"))) {
//             scanningHeaders = false;
//             scanningBody = true;
//             continue;
//           }
//         }

//         if (scanningBody) {
//           body += line;
//         }

//         if ((line.length() == 1 && line[0] == '\n')) {
//           // if (fn != NULL) {
//           //   response *r = (fn)(headers, body);

//           //   if (r->status == 200) {
//           //     client.print("HTTP/1.1 200 OK\r\nContent-Type: "
//           //                  "text/html\r\nConnection: close\r\n\r\n" +
//           //                  r->body);
//           //   } else {
//           //     client.print("HTTP/1.1 500 OK\r\nContent-Type: "
//           //                  "text/html\r\nConnection: close\r\n\r\n");
//           //   }
//           // } else {
//           //   client.print("HTTP/1.0 404 Not Found\r\nContent-Type: "
//           //                "text/html\r\nConnection: close\r\n\r\n");
//           // }
//           client.println("HTTP/1.1 200 OK\r\nContent-Type: "
//                          "text/html\r\nConnection: close\r\n");
//           break;
//         }
//       }
//     }
//     delay(1);
//     client.stop();
//   }
// }