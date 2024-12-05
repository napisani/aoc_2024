import gleam/int
import gleam/io
import gleam/list
import gleam/option.{Some}
import gleam/regexp
import gleam/result
import simplifile

pub fn main() {
  let assert Ok(contents) = simplifile.read("./input/day3.txt")
  let contents = "do()" <> contents
  let dont_regex =
    result.lazy_unwrap(regexp.from_string("don't\\(\\)"), fn() {
      panic as "Invalid regexp"
    })
  let do_regex =
    result.lazy_unwrap(regexp.from_string("do\\(\\)"), fn() {
      panic as "Invalid regexp"
    })
  let mul_regex =
    result.lazy_unwrap(
      regexp.from_string("mul\\(([\\d]{1,3}),([\\d]{1,3})\\)"),
      fn() { panic as "Invalid mul regexp" },
    )
  let split_by_dont = regexp.split(dont_regex, contents)

  let final =
    list.fold(split_by_dont, 0, fn(f, contents) {
      let split_by_do = regexp.split(do_regex, contents)

      f
      + list.index_fold(split_by_do, 0, fn(total, contents, idx) {
        case idx > 0 {
          False -> {
            total
          }
          True -> {
            // io.debug(contents <> " processed")
            let matches = regexp.scan(mul_regex, contents)
            total
            + list.fold(matches, 0, fn(acc, m) {
              case m.submatches {
                [Some(a), Some(b)] -> {
                  let a_int = result.unwrap(int.parse(a), 0)
                  let b_int = result.unwrap(int.parse(b), 0)
                  io.debug(m)
                  acc + a_int * b_int
                }
                _ -> acc
              }
            })
          }
        }
      })
    })
  io.debug(final)
}
