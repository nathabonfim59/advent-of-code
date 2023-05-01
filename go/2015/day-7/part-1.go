package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
    OP_AND = "AND"
    OP_OR = "OR"
    OP_NOT = "NOT"
    OP_LSHIFT = "LSHIFT"
    OP_RSHIFT = "RSHIFT"
    OP_ASSIGN = "ASSIGN"
)

type Wire struct {
    identifier string
    signal uint16
}

// The gate
type Instruction struct {
    operation string
    shift uint16
    inputs []Wire
    output Wire
}

func main() {
    var inputFile = "input.data"
    var fileReader *bufio.Reader = readFile(inputFile)
    var rawInstruction string
    var instruction Instruction
    var circuit []Wire

    for {
        rawInstruction = getNextString(fileReader)
        if rawInstruction == "" {
            break
        }

        instruction = parseInstruction(rawInstruction)
        processInstruction(instruction, &circuit)

        fmt.Println("----")
        fmt.Println("Instruction: ", instruction)
        fmt.Println("Circuit: ", circuit)
    }

    // printCircuit(&circuit)
}

// Section: Input processing
//--------------------------------

func readFile(inputFile string) *bufio.Reader {
    file, err := os.Open(inputFile)

    if err != nil {
        fmt.Println("Error reading the file", err)
        os.Exit(1)
    }

    return bufio.NewReader(file)
}

func getNextString(fileReader *bufio.Reader) string {
    line, err := fileReader.ReadString('\n')

    // If it is the end of the file, return an empty string
    if err != nil {
        return ""
    }

    // Removes the newline
    line = line[:len(line) - 1]

    return line
}

// Section: Parsing
//--------------------------------

func parseInstruction(rawInstruction string) Instruction {
    // Example: x LSHIFT 2 -> f

    var instruction Instruction
    var instructionParts []string = strings.Split(rawInstruction, "->")
    var rawInput []string = strings.Split(string(instructionParts[0]), " ")
    var rawOutput []string = strings.Split(string(instructionParts[1]), " ")

    // Remove the trailing space
    rawInput = rawInput[:len(rawInput) - 1]

    // Remove the leading space
    rawOutput = rawOutput[1:]

    instruction.operation = parseOperation(rawInput)
    instruction = parseInputs(rawInput, instruction)
    instruction.output = parseOutput(rawOutput)

    return instruction
}

func parseOperation(rawInput []string) string {
    var rawOperation string
    var operation string
    

    var isAssignInput bool = len(rawInput) == 1; // input
    var isNotInput bool = len(rawInput) == 2;    // operation + input
    var isOtherOpInput = len(rawInput) == 3;     // input1 + operation + input2

    if (isAssignInput) {
        rawOperation = OP_ASSIGN
    } else if (isNotInput) {
        rawOperation = OP_NOT
    } else if (isOtherOpInput) {
        rawOperation = rawInput[1]
    }

    switch rawOperation {
        case OP_AND:
            operation = OP_AND
        case OP_OR:
            operation = OP_OR
        case OP_NOT:
            operation = OP_NOT
        case OP_LSHIFT:
            operation = OP_LSHIFT
        case OP_RSHIFT:
            operation = OP_RSHIFT
        case OP_ASSIGN:
            operation = OP_ASSIGN
    }

    return operation
}

func parseInputs(rawInput []string, instruction Instruction) Instruction {
    var inputs []Wire

    switch instruction.operation {
        case OP_ASSIGN:
            signal, _ := strconv.Atoi(rawInput[0])
            inputs = append(inputs, Wire{"", uint16(signal)})

        case OP_NOT:
            identifier := rawInput[1]
            inputs = append(inputs, Wire{identifier, 0})

        case OP_AND, OP_OR:
            identifier_1 := rawInput[0]
            identifier_2 := rawInput[2]

            // If the identifier is a signal, add it as a signal
            if _, err := strconv.Atoi(identifier_1); err == nil {
                signal, _ := strconv.Atoi(identifier_1)
                inputs = append(inputs, Wire{"", uint16(signal)})
            } else {
                inputs = append(inputs, Wire{identifier_1, 0})
            }

            if _, err := strconv.Atoi(identifier_2); err == nil {
                signal, _ := strconv.Atoi(identifier_2)
                inputs = append(inputs, Wire{"", uint16(signal)})
            } else {
                inputs = append(inputs, Wire{identifier_2, 0})
            }

        case OP_LSHIFT, OP_RSHIFT:
            identifier := rawInput[0]
            shift, _ := strconv.Atoi(rawInput[2])

            inputs = append(inputs, Wire{identifier, 0})
            instruction.shift = uint16(shift)
    }

    instruction.inputs = inputs

    return instruction
}

func parseOutput(rawOutput []string) Wire {
    var identifier string = rawOutput[0]

    return Wire{identifier, 0}
}

func processInstruction(instruction Instruction, circuit *[]Wire) {
    switch instruction.operation {
        case OP_ASSIGN:
            assign(instruction, circuit)

        case OP_NOT:
            not(instruction, circuit)

        case OP_AND:
            and(instruction, circuit)

        case OP_OR:
            or(instruction, circuit)

        case OP_LSHIFT:
            lshift(instruction, circuit)

        case OP_RSHIFT:
            rshift(instruction, circuit)
    }
}

func getWire(targetWire Wire, circuit *[]Wire) Wire {
    for _, wire := range *circuit {
        if wire.identifier == targetWire.identifier {
            return wire
        }
    }

    return targetWire
}

func replaceOrCreateWire(wire Wire, circuit *[]Wire) {
    for index, circuitWire := range *circuit {
        if circuitWire.identifier == wire.identifier {
            (*circuit)[index] = wire
            return
        }
    }

    (*circuit) = append(*circuit, wire)
}

func sortCircuit(circuit *[]Wire) {
    sort.Slice(*circuit, func(i, j int) bool {
        if (*circuit)[i].identifier < (*circuit)[j].identifier {
            return true
        }
        return false
    })
}

func printCircuit(circuit *[]Wire) {
    // Print the circuit in alphabetical ordeer
    sortCircuit(circuit)

    for _, wire := range *circuit {
        fmt.Println(wire.identifier, ": ", wire.signal)
    }
}

// Section: Bitwise operations
//--------------------------------


func assign(instruction Instruction, circuit *[]Wire) {
    var outputWire Wire

    // Verify if the wire already exists
    outputWire = getWire(instruction.output, circuit)

    if outputWire.identifier == "" {
        outputWire = Wire{
            instruction.output.identifier,
            instruction.inputs[0].signal,
        }
    } else {
        outputWire.signal = instruction.inputs[0].signal
    }

    replaceOrCreateWire(outputWire, circuit)
}

func not(instruction Instruction, circuit *[]Wire) {
    var outputWire Wire
    var inputWire Wire

    outputWire.identifier = instruction.output.identifier
    inputWire = getWire(instruction.inputs[0], circuit)

    outputWire.signal = uint16(^inputWire.signal)

    replaceOrCreateWire(outputWire, circuit)
}

func and(instruction Instruction, circuit *[]Wire) {
    var outputWire Wire
    var input1 Wire = getWire(instruction.inputs[0], circuit)
    var input2 Wire = getWire(instruction.inputs[1], circuit)

    outputWire.identifier = instruction.output.identifier
    outputWire.signal = input1.signal & input2.signal

    replaceOrCreateWire(outputWire, circuit)
}

func or(instruction Instruction, circuit *[]Wire) {
    var outputWire Wire
    var input1 Wire = getWire(instruction.inputs[0], circuit)
    var input2 Wire = getWire(instruction.inputs[1], circuit)

    outputWire.identifier = instruction.output.identifier
    outputWire.signal = input1.signal | input2.signal

    replaceOrCreateWire(outputWire, circuit)
}

func lshift(instruction Instruction, circuit *[]Wire) {
    var outputWire Wire
    var inputWire Wire = getWire(instruction.inputs[0], circuit)

    outputWire.identifier = instruction.output.identifier
    outputWire.signal = inputWire.signal << instruction.shift

    replaceOrCreateWire(outputWire, circuit)
}

func rshift(instruction Instruction, circuit *[]Wire) {
    var outputWire Wire
    var inputWire Wire = getWire(instruction.inputs[0], circuit)

    outputWire.identifier = instruction.output.identifier
    outputWire.signal = inputWire.signal >> instruction.shift

    replaceOrCreateWire(outputWire, circuit)
}


