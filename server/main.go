package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Guide
// ./usercli  --help
// ./usercli submit --help

// 실행방법
// 1. Run 'go build -o `filename` main.go'
// 2. Run './`filename` submit -i `input`'

func main() {
	var rootCmd = &cobra.Command{Use: "usercli"} // Cobra를 사용하여 CLI의 루트 명령을 설정합니다.
	rootCmd.AddCommand(submitIDCmd())            // submitIDCmd 함수로부터 반환된 명령을 루트 명령에 추가합니다.
	if err := rootCmd.Execute(); err != nil {    // CLI 애플리케이션을 실행하고 오류를 처리합니다.
		fmt.Println(err)
		os.Exit(1)
	}
}

func submitIDCmd() *cobra.Command { // 사용자 ID를 제출하기 위한 Cobra 명령을 정의하는 함수입니다.
	var userID string // 사용자 ID를 저장할 변수를 선언합니다.
	var cmd = &cobra.Command{
		Use:   "submit",                                                       // 명령의 이름을 정의합니다.
		Short: "Submit user ID to the blockchain",                             // 명령의 간단한 설명을 제공합니다.
		Long:  `This command submits a user ID to a blockchain-based system.`, // 명령의 긴 설명을 제공합니다.
		Run: func(cmd *cobra.Command, args []string) { // 명령이 호출될 때 실행될 함수를 정의합니다.
			fmt.Printf("Processing ID: %s\n", userID)                        // 처리 중인 ID를 출력합니다.
			if success, err := storeUserIDInBlockchain(userID); err != nil { // Blockchain에 ID를 저장하는 함수를 호출하고 오류를 처리합니다.
				fmt.Printf("Failed to store ID: %v\n", err) // ID 저장 실패 시 오류 메시지를 출력합니다.
			} else {
				fmt.Println("ID successfully stored in the blockchain:", success) // ID 저장 성공 메시지를 출력합니다.
			}
		},
	}
	cmd.Flags().StringVarP(&userID, "id", "i", "", "User ID to submit") // 명령에 userID 입력을 위한 플래그를 추가합니다.
	cmd.MarkFlagRequired("id")                                          // 'id' 플래그를 필수로 설정합니다.
	return cmd
}

// This function simulates interaction with a smart contract. // 스마트 컨트랙트와 상호작용을 시뮬레이션하는 함수입니다.
func storeUserIDInBlockchain(userID string) (bool, error) { // 사용자 ID를 블록체인에 저장합니다.
	// Implement the logic to interact with the smart contract. // 스마트 컨트랙트와 상호작용하는 로직을 구현해야 합니다.
	// Assuming the operation is successful. // 작업이 성공적이라고 가정합니다.
	return true, nil // 성공적으로 저장됐다고 반환합니다.
}
