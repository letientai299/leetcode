use divan::AllocProfiler;
use lc_rust::util::*;

#[global_allocator]
static ALLOC: AllocProfiler = AllocProfiler::system();

fn main() {
    divan::main();
}

// Register a `fibonacci` function and benchmark it over multiple cases.
#[divan::bench]
fn p_2657_find_the_prefix_common_array() -> Vec<i32> {
    use lc_rust::p_2657_find_the_prefix_common_array::Solution;

    let a = vec_of::<i32>("[1,3,2,4]");
    let b = vec_of::<i32>("[3,1,2,4]");

    Solution::find_the_prefix_common_array(
        divan::black_box(a),
        divan::black_box(b),
    )
}
