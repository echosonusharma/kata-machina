# KATA-MACHINA

A cli tool to practice DSA (in go)

The `dsa.config.json` file contains all the configuration for generating the DSA you want to practice.

- DSA names added under the `dsa` arr in the config will generate those dsa for you to practice.

**List of supported DSA**

- `linearSearchList`
- `BinarySearchList`
- `TwoCrystalBalls`
- `BubbleSort`
- `SinglyLinkedList`
- `Queue`
- `Stack`
- `MazeSolver`

**Commands for the CLI**

- `gen` - generates a new dir with a list of dsa files with their starter code. 
- `clean` - removes last created dsa dir or you can pass arg -all to remove all generated dsa dir.
- `test` - by default it will run the test for the latest generated dsa dir created or you can pass the dir name for which you want to run the tests for.