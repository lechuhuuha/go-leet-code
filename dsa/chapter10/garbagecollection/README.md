# Introductions

- Garbage collection is a type of programmed memory management in which memory, currently occupied by objects that will never be used again, is gathered.

- It's a system that recycles memory on behalf of the application by indentifying which parts of memory are no longer needed.

- John McCarthy was the first person to come up with garbage collection for managing Lisp memory management.

- This technique specifies which objects need to be de-allocated, and then discharges the memory.

- **Socket, database connections, user window objects and file handle** are not overseen by garbage collectors.

- Understand costs by looking into the Go GC.

- Use these insights to improve application's resource utilization.

- The Go standard toolchain provides a runtime library that ships with every application, and this runtime library includes a garbage collector.

### Terms:

- **Object**

  - An object is dynamically allocated piece of memory that contains one or more Go values.

- **Pointer**
  - A memory address that references any value within an object. This includes Go values of the form \*T, but also includes parts of built-in Go value. Strings, slices, channels, maps and interface values. All contain memory address that GC **must trace.**

# Where Go values live

- Non-pointer Go values stored in local variables will likely **not be managed by the Go GC at all**

  - Go will instead arrange for memory to be allocated that's tied to the _lexical scope_ in which it's created.

    - This mean that when a go variable is created in function scope. Then its will belong to that func stack. Once the func exits that value will be remove automatically with the func stack

  - In genaral, this is more efficient then reyling on the GC, because the Go compiler is able to predetermine when that memory maybe freed and emit machine instructions that clean it up.

  - Typically, allocating memory for Go values this way as **"stack allocation"**, because the space is stored on the goroutine stack.

- Go values whose memory cannot be allocated this way, because Go compiler cannot determine its lifetime, are said to escape to the heap.

- **"The heap"** can be thought of as a catch-all for memory allocation, for when Go values need to place somewhere.

- The act of allocation memory on the heap is typically referred to as **"dynamic memory alloction"** because both the compiler and the runtime can make very few assumptions as to how this memory is used and when it can be cleaned up.

- That's where a GC comes in: it's a system that specifically **identifies and cleans up dynamic memory allocations.**

- Some reasons why a Go value might need to escape to the heap.

  - The size is **dynamically determined.**
  - Array of slice whose initial size is determined by a variable, rather then a constant.
  - Escaping to the heap must also be **transitive**.
    - If a reference to a Go value is written into another Go value that has **already been determined to escape**, that value must also escape.

# Tracing Garbage collection

- GC may refer to many different method of automatically recyling memory; for example, **reference counting.**

  - Reference counting is related to **keeping the number of references**, pointers, and handles to a resource such as an object, block of memory, or disk space.
  - This technique is related to the number of references to de-allocated objects that are never referenced again.
  - The number of references is used for runtime optimizations. **Deutsch-Bobrow** came up with the strategy of reference counting.
  - This strat was related to the number of updated references that were produced by references that were put in local variables.
  - **Henry Baker** came up with a method that includes references in local variables that are deferred until needed.

- In Go, GC refer to **tracing** garbage collection, which identifies in-use, so-called **live** object by following pointers transitively.

- Together, objects and pointers to other objects form the **object graph**

- To identify live memory, the GC walk the object graph starting at the program's **roots**, pointers that identify objects that are currently in-use by the program.

- **Examples:**

  - Local vars
  - Global vars

- The process of walking the object graph is referred to as **Scanning.**

- This basic algorithm is **common to all tracing GCs**. Where tracing GCs differ is **what they do once they discover memory is live.**

- Go's GC uses the mark-sweep technique, which means that in order to keep track of its progress, the GC also **marks** the values it encounters as live.

  - Once tracing is complete, the GC then walks over all memory in the heap and makes all memory that is _not_ marked available for allocation. This process is called sweeping.

- Go has a **non-moving GC**
  - a garbage collector that **cannot move live data**. This is different from **a moving GC**, which **can move live data** to free up consecutive locations for free space.

# The GC cycle

- Operates in two phases:

  - the mark phase
  - the sweep phase

- Important notes:

  - It's not possible to release memory back to be allocated until all memory has been traced, because there may still be an un-scanned pointer keeping an object alive.

  * As a result, the act of sweeping must be entirely separated from the act of marking. Furthermore, the GC may also not be active at all, when there's no GC-related work to do.
  * The GC continuously rotates through three phase of sweeping, off and marking in what's known as the GC cycle.

## Understanding costs

- The GC is a complex piece of software built on even more complex systems.
- The model of GC cost based on three simple axioms.
  - The GC involves only two resources: CPU time and physical memory.
  - The GC's memory costs consist of live heap memory, new heap memory allocated before the mark phase, and space for metadata that, even if proportional to the previous costs, are small in comparison.
    _Note: live heap memory is memory that was determined to be live by the previous GC cycle, while new heap memory is any memory allocateed in the current cycle, which may or may not be live by the end._
  - The GC's CPU costs are modeled as a fixed cost per cycle, and marginal cost that scales proportionally with the size of the live heap.
    _Note: Asymptotically speaking, sweeping scales worse than marking and scanning, as it must perform work proportional to the size of the whole heap, including memory that is determined to be not live ("dead"), in the current implentaion sweeping is so much faster then marking and scanning that its associated costs can be ignored in this discussion._
