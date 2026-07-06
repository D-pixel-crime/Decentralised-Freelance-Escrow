package main

import (
	"os"
	"os/exec"

	// Indexer disabled — JIT sync in GetMyJobs replaces push-based WebSocket listeners
	// "github.com/D-pixel-crime/Freelance_Escrow/backend/indexer"
	authRouter "github.com/D-pixel-crime/Freelance_Escrow/backend/routers/authRouters"
	getRouter "github.com/D-pixel-crime/Freelance_Escrow/backend/routers/getRouters"
	postRouter "github.com/D-pixel-crime/Freelance_Escrow/backend/routers/postRouters"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/gin-contrib/cors"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func generateBindings() {
	log.Info("Compiling smart contracts and generating Go bindings...")

	// Step 1: Run forge build
	buildCmd := exec.Command("forge", "build", "--extra-output-files", "abi", "bin")
	buildCmd.Dir = "../contracts"
	buildOut, err := buildCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run forge build: %v\nOutput: %s", err, string(buildOut))
	}

	// Step 2: Run abigen
	abigenCmd := exec.Command("abigen",
		"--abi", "../contracts/out/FreelanceEscrow.sol/FreelanceEscrow.abi.json",
		"--bin", "../contracts/out/FreelanceEscrow.sol/FreelanceEscrow.bin",
		"--pkg", "contracts",
		"--type", "FreelanceEscrow",
		"--out", "./contracts/FreelanceEscrow.go",
	)
	abiOut, err := abigenCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run abigen: %v\nOutput: %s", err, string(abiOut))
	}

	log.Info("Successfully generated smart contract bindings!")
}

func main() {
	generateBindings()
	utils.LoadENV()
	_, err := utils.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	_, err = utils.ConnectToRedis()
	if err != nil {
		log.Fatal(err)
	}

	_, err = utils.ConnectToWeb3()
	if err != nil {
		log.Errorf("Failed to connect to Web3: %v", err)
	}
	// Phase 11.2: Old push-based indexer disabled.
	// On-chain state is now reconciled JIT (Just-In-Time) inside the GetMyJobs handler.

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	authRouter.AUTH_Routes(router)
	getRouter.GET_Routes(router)
	postRouter.POST_Routes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Infof("Server starting on port: %s...", port)
	router.Run(":" + port)
}
