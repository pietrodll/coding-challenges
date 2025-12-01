import day15
import gleam/pair
import gleeunit/should
import utils_matrix

const input1 = "##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^"

const input2 = "########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<"

pub fn move_robot_once_test() {
  let map =
    "########
#.@O.O.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########"
    |> day15.parse_map
    |> should.be_ok

  let expected =
    "########
#..@OO.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########"
    |> day15.parse_map
    |> should.be_ok

  day15.move_robot_once(map, #(1, 2), utils_matrix.Right)
  |> pair.first
  |> should.equal(expected)
}

pub fn move_robot_test() {
  let #(map1, directions1) = input1 |> day15.parse_input |> should.be_ok

  let expected1 =
    "##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########"
    |> day15.parse_map
    |> should.be_ok

  day15.move_robot(map1, directions1) |> should.equal(expected1)

  let #(map2, directions2) = input2 |> day15.parse_input |> should.be_ok

  let expected2 =
    "########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########"
    |> day15.parse_map
    |> should.be_ok

  day15.move_robot(map2, directions2) |> should.equal(expected2)
}
