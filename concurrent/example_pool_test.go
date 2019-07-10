package concurrent_test

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/fblaha/manaus-export-import/concurrent"
	"hash"
	"io"
	"io/ioutil"
	"os"
)

// hashInput contains all inputs needed to hash single file
type hashInput struct {
	// file name
	file string
	// hash factory
	hashFactory func() hash.Hash
}

// hashOutput contains the hashing output or error
type hashOutput struct {
	hashInput
	value []byte
	err   error
}

// hashWork hashes single file according to inputs
type hashWork struct {
	hashInput
	out chan<- hashOutput
}

// Work implements Worker interface
// perform hashing and result writes to output channel
func (hw hashWork) Work() {
	output := hashOutput{hashInput: hw.hashInput}
	f, err := os.Open(hw.file)
	if err != nil {
		// propagates error to output and returns
		output.err = err
		hw.out <- output
		return
	}
	defer f.Close()

	// creates an instance of hash
	h := hw.hashFactory()
	if _, err := io.Copy(h, f); err != nil {
		// propagates error to output and returns
		output.err = err
		hw.out <- output
		return
	}
	// writes computed hash to output channel
	output.value = h.Sum(nil)
	hw.out <- output
}

// Example demonstrates usage of pool executor with graceful shutdown and result/error propagation,
// The example iterates all files in current directory.
// For each file computes sha256, sha384 and sha512 hash
func Example() {
	// output channel
	output := make(chan hashOutput)
	// slices with hashes which we want to compute for each file
	hashFactories := []func() hash.Hash{sha256.New, sha512.New384, sha512.New}

	go func() {
		// constructs the executor
		executor := concurrent.NewPoolExecutor(10)
		// read files from current dir
		files, _ := ioutil.ReadDir(".")
		for _, file := range files {
			for _, f := range hashFactories {
				// submit hash work for execution
				// for each file and each hash
				executor.Submit(hashWork{out: output, hashInput: hashInput{file: file.Name(), hashFactory: f}})
			}
		}
		// wait for completion and shutdowns executor
		executor.ShutdownGracefully()
		// closes output channel
		close(output)

	}()

	// iterates output channel and prints results
	for result := range output {
		if result.err != nil {
			fmt.Printf("hashing of %s failed: %+v\n", result.file, result.err)
			continue
		}
		fmt.Printf("%s (size: %d bytes) : %x \n", result.file,result.hashFactory().Size(), result.value)
	}
}

//OUTPUT:
//example_pool_test.go (size: 32 bytes) : 41f97bfde600c84515184cec2afe2e145fb5ab3babaae0492693151895e1707e
//pool_test.go (size: 64 bytes) : 08ab4652993b874a5de215067ff82fb4acfbf793fae3d5f229268bfac7c00fc52f43a73eacad9f1d8954fa1b95b64daf760dee30c14fc3d606c1e3d8e09fcefd
//example_pool_test.go (size: 64 bytes) : 2a1bd7c03df0d49856f5d2d825f8331156f6d37457ed916d23e3d009454f6a8bcc8b755d8bdf58ad037e82651e3570ade356a40c1e2d6f1b1f26b8aafac46021
//pool_test.go (size: 32 bytes) : 37ca25eb57fc6590349b42e4ab4883451aa73ee3bf171c0827707cdc0c621de1
//example_pool_test.go (size: 48 bytes) : 626544c44777e4b55f10c4c09f3fe1580b1af52173195a633d0b1f6a3505ea99e43a0d521d29cafe6cd740dfe49fafed
//pool.go (size: 64 bytes) : 1eb1ee8aebe2bae793848e7316ea9eeca477fe09a248bf9bbaeed72fbc45515bf8c2f17644b48d9dd88275d8a61aae07091f321fa66542f2c95f0f84c10a96ad
//pool_test.go (size: 48 bytes) : 60a87b791dbf8bb9bc16154ec35f3c6caff2b701d9f5e4da4cb2e1089ccd419ff9591ece431f81a72bcd4899cac7737d
//pool.go (size: 32 bytes) : 50f9e39091e96baef7e2c727b5631792634d5eaa0c7d909c741b02f017dca6aa
//pool.go (size: 48 bytes) : 96a8b8d12c71a55266b9219cc4dc252e9ecc9de43031b6aa62dad0589883d667556e80e52b548c25bc238f5dc816d502
//
//func Test(t *testing.T) {
//	Example()
//}
