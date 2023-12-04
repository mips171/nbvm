package main

import "fmt"

type VM struct {
    registers [4]int
    memory    []int
    pc        int
}

// Instruction set
const (
    HALT  = iota
    LOAD
    ADD
    PRINT
)

func (vm *VM) Run() {
    for {
        switch vm.memory[vm.pc] {
        case HALT:
            return
        case LOAD:
            register := vm.memory[vm.pc+1]
            value := vm.memory[vm.pc+2]
            vm.registers[register] = value
            vm.pc += 3
        case ADD:
            reg1 := vm.memory[vm.pc+1]
            reg2 := vm.memory[vm.pc+2]
            vm.registers[reg1] += vm.registers[reg2]
            vm.pc += 3
        case PRINT:
            fmt.Println(vm.registers[vm.memory[vm.pc+1]])
            vm.pc += 2
        default:
            fmt.Println("Unknown instruction")
            return
        }
    }
}

func main() {
    vm := VM{
        memory: []int{
            LOAD, 0, 10,
            LOAD, 1, 20,
            ADD, 0, 1,
            PRINT, 0,
            HALT,
        },
    }

    vm.Run()
}
