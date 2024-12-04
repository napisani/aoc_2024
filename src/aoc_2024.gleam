import gleam/io
import gleam/regexp
import simplifile
import gleam/list
import gleam/int
import gleam/option.{Some}
import gleam/result


pub fn main() {
  let contents = result.lazy_unwrap(simplifile.read("./input/day3.txt"), fn () {panic as "Could not read file"})
  let contents = "do()" <> contents
  let dont_regex = result.lazy_unwrap(regexp.from_string("don't\\(\\)"), fn () {panic as "Invalid regexp"})
  let do_regex = result.lazy_unwrap(regexp.from_string("do\\(\\)"), fn () {panic as "Invalid regexp"})
  let mul_regex = result.lazy_unwrap(regexp.from_string("mul\\(([\\d]{1,3}),([\\d]{1,3})\\)"), fn () {panic as "Invalid mul regexp"})
  let split_by_dont = regexp.split(dont_regex, contents)

  let final = list.fold(split_by_dont, 0, fn (f, contents) {
    let split_by_do = regexp.split(do_regex, contents)

      f + list.index_fold(split_by_do, 0, fn ( total,contents,  idx) {
        io.debug(idx)
        case idx > 0  {
          False -> {
          io.debug(contents <> " NOT processed")
          total
        }
          True -> {
            io.debug(contents <> " processed")
            let matches = regexp.scan(mul_regex, contents) 
            total + list.fold(matches,0,  fn (acc, m) { 
              case m.submatches {
                [Some(a), Some(b)] -> {
                  let a_int = result.unwrap(int.parse(a), 0)
                  let b_int =result.unwrap(int.parse(b), 0)
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
