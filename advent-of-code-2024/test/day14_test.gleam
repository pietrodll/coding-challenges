import day14
import gleeunit/should

const input = "p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3"

const grid_height = 7

const grid_width = 11

pub fn move_test() {
  let robot = day14.Robot(#(2, 4), #(2, -3))

  day14.move(robot, 0, grid_height, grid_width) |> should.equal(#(2, 4))
  day14.move(robot, 1, grid_height, grid_width) |> should.equal(#(4, 1))
  day14.move(robot, 2, grid_height, grid_width) |> should.equal(#(6, 5))
}

pub fn quadrant_test() {
  day14.quadrant(6, 0, grid_height, grid_width) |> should.equal(day14.TopRight)

  day14.quadrant(0, 2, grid_height, grid_width) |> should.equal(day14.TopLeft)

  day14.quadrant(3, 5, grid_height, grid_width)
  |> should.equal(day14.BottomLeft)

  day14.quadrant(6, 6, grid_height, grid_width)
  |> should.equal(day14.BottomRight)
}

pub fn safety_factor_test() {
  input
  |> day14.parse_input
  |> should.be_ok
  |> day14.safety_factor(100, grid_height, grid_width)
  |> should.equal(12)
}
