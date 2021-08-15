package glfw

//go:generate ../../scripts/glfw_tree_rebuild.sh

// upstreamTreeSHA is a recursive hash of the full contents of the upstream
// glfw, as generated by git (doesn't need to be committed) when you run `go
// generate` on this package. This exists to invalidate the build cache (see
// https://github.com/go-gl/glfw/issues/269), which is unaffected by C source
// inputs.
const upstreamTreeSHA = "7e7e4811fbddc9425712743eea77cb1668528e76"
