<img src="adapter-mini-2x.png" alt="adapter-mini-2x" style="width: 150px; height: 100px;" /> <span style="font-size: 50px;">Adapter in Go</span>

**Adapter** is a structural design pattern, which allows incompatible objects to collaborate.

The Adapter acts as a wrapper between two objects. It catches calls for one object and transforms them to format and interface recognizable by the second object.

# Conceptual Example

We have a client code that expects some features of an object (Lightning port), but we have another object called adaptee (Windows laptop) which offers the same functionality but through a different interface (USB port)

This is where the Adapter pattern comes into the picture. We create a struct type known as adapter that will:

- Adhere to the same interface which the client expects (Lightning port).
- Translate the request from the client to the adaptee in the form that the adaptee expects.
- The adapter accepts a Lightning connector and:
  - Translates its signals into a USB format.
  - Passes the signals to the USB port in a Windows laptop.
