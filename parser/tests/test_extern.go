package main

func init() {
	RegisterTest("Extern Declaration", func() string {
		return `extern os {
    func printf(*byte, ...) int32
    func sleep(int32) int32
    func usleep(int32) int32
}`
	})

	RegisterTest("Extern With Alias", func() string {
		return `extern libc {
    func printf "printf" (*byte, ...) int32
    func sleep "usleep" (int32) int32
}`
	})

	RegisterTest("Extern Direct Mapping", func() string {
		return `extern libc {
    func usleep(int32) int32
}`
	})
}