## Decorator pattern

The decorator pattern attaches additional responsibilities to an object dynamically. 
The decorators provide a flexible alternative to subclassing for extending functionality.

## Notes
Go is not an Object Oriented language. We can take advantage of the way the interfaces work in Go: to implement it we just need to implement all the methods that are declared in the interface.

In the example, using beverages, any variable will be of the "class" beverage if it implements the methods in the beverage interface. Thus, there's no need to declare inheritance of classes or to have a superclass.