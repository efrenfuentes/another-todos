use actix_web::{delete, get, post, put, web, Responder};
use actix_web::http::StatusCode;

use crate::models::todo::{TodoPayload};
use crate::models::todo;
use crate::DbPool;
use crate::helpers;

#[get("")]
async fn index(pool: web::Data<DbPool>) -> impl Responder {
    let mut conn = pool.get().unwrap();
    let results = todo::all(&mut conn).unwrap();

    helpers::response::json_response(StatusCode::OK, "todos", results)
}

#[get("/{id}")]
async fn show(path: web::Path<i32>, pool: web::Data<DbPool>) -> impl Responder {
    let mut conn = pool.get().unwrap();

    let todo_id = path.into_inner();

    let todo = todo::find(todo_id, &mut conn);

    match todo {
        Ok(todo) => helpers::response::json_response(StatusCode::OK, "todo", todo),
        Err(e) => helpers::response::json_error_db_response(e)
    }
}

#[post("")]
async fn create(pool: web::Data<DbPool>, payload: web::Json<TodoPayload>) -> impl Responder {
    let mut conn = pool.get().unwrap();
    let new_todo = todo::add(payload.into_inner(), &mut conn);

    match new_todo {
        Ok(todo) => helpers::response::json_response(StatusCode::OK, "todo", todo),
        Err(e) => helpers::response::json_error_db_response(e)
    }
}

#[put("/{id}")]
async fn update(path: web::Path<i32>, pool: web::Data<DbPool>, payload: web::Json<TodoPayload>) -> impl Responder {
    let mut conn = pool.get().unwrap();

    let todo_id = path.into_inner();

    let updated_todo = todo::update(todo_id, payload.into_inner(), &mut conn);

    match updated_todo {
        Ok(todo) => helpers::response::json_response(StatusCode::OK, "todo", todo),
        Err(e) => helpers::response::json_error_db_response(e)
    }
}

#[delete("/{id}")]
async fn delete(path: web::Path<i32>, pool: web::Data<DbPool>) -> impl Responder {
    let mut conn = pool.get().unwrap();

    let todo_id = path.into_inner();

    let todo = todo::delete(todo_id, &mut conn);

    match todo {
        Ok(todo) => helpers::response::json_response(StatusCode::OK, "todo", todo),
        Err(e) => helpers::response::json_error_db_response(e)
    }
}
