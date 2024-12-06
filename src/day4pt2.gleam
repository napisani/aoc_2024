import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/option.{type Option, None, Some}
import gleam/result
import gleam/string
import gleam/yielder
import simplifile

const word = "MAS"

pub type Match {
  Match(coord: #(Int, Int), direction: #(Int, Int))
}

fn get_match(
  d: Dict(#(Int, Int), String),
  coord: #(Int, Int),
  dir: #(Int, Int),
  y_len: Int,
  x_len: Int,
) -> Option(Match) {
  let assert Ok(c) = dict.get(d, coord)
  let word_len = string.length(word) - 1

  let furthest_y = coord.0 + { dir.0 * word_len }
  let furthest_x = coord.1 + { dir.1 * word_len }

  let has_room =
    furthest_y >= 0
    && furthest_x >= 0
    && furthest_y < y_len
    && furthest_x < x_len
    && c == "M"

  case has_room {
    False -> None
    True -> {
      let found =
        iterator.range(from: 1, to: word_len)
        |> iterator.to_list()
        |> list.fold("M", fn(acc, i) {
          let next_coord = #(coord.0 + { dir.0 * i }, coord.1 + { dir.1 * i })
          let assert Ok(next_c) = dict.get(d, next_coord)
          acc <> next_c
        })
      case found == word {
        True -> {
          io.debug(found)
          io.debug(#(coord.0 + dir.0, coord.1 + dir.1))
          Some(Match(coord: #(coord.0 + dir.0, coord.1 + dir.1), direction: dir))
        }
        False -> None
      }
    }
  }
}

pub fn main() {
  let assert Ok(contents) = simplifile.read("./input/day4.txt")
  let directions = [#(-1, -1), #(1, 1), #(-1, 1), #(1, -1)]
  let split = string.split(contents, "\n") |> list.filter(fn(s) { s != "" })

  let content_map =
    list.index_fold(split, dict.new(), fn(acc, line, y) {
      let line_list = string.split(line, "")
      list.index_fold(line_list, acc, fn(acc, c, x) {
        dict.insert(acc, #(y, x), c)
      })
    })

  let y_len = list.length(split)
  let assert Ok(first) = list.first(split)
  let x_len = string.length(first)

  let matches =
    dict.keys(content_map)
    |> list.flat_map(fn(coord) {
      list.map(directions, fn(dir) {
        get_match(content_map, coord, dir, y_len, x_len)
      })
    })
    |> list.filter_map(fn(m) {
      case m {
        Some(m) -> Ok(m)
        None -> Error(Nil)
      }
    })

  let unique =
    matches
    |> list.fold(dict.new(), fn(acc, m) { dict.insert(acc, m.coord, m) })

  let unique_list = dict.keys(unique)
  let unique_list =
    list.map(unique_list, fn(coord) {
      let x = int.to_string(coord.0) <> " " <> int.to_string(coord.1)
      io.debug(x)
      x
    })
  io.debug(list.length(unique_list))
}
