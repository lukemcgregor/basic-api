#![feature(plugin)]
#![plugin(rocket_codegen)]

extern crate rocket;
extern crate elementtree;
#[macro_use] extern crate rocket_contrib;

use rocket::data::Data;
use rocket_contrib::{JSON, Value};
use elementtree::Element;

#[get("/add/<x>/to/<y>")]
fn add_json(x: i32, y: i32) -> JSON<Value> {
    JSON(json!({ "sum": (x+y) }))
}

#[post("/add", data = "<numbers>")]
fn add_xml(numbers: Data) -> String {
	let root = Element::from_reader(numbers.open()).unwrap();
	let sum : i32 = root.find_all("value").map(|e| e.text().parse::<i32>().unwrap()).sum();

	let mut result = Element::new("sum");
	result.set_text(sum.to_string());

	return result.to_string().unwrap();
}

fn main() {
    rocket::ignite().mount("/", routes![add_json, add_xml]).launch();
}