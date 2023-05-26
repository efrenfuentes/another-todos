use serde::{Deserialize, Serialize};

use crate::models::schema::todos;
use crate::models::errors::DbError;
use diesel::PgConnection;
use diesel::RunQueryDsl;
use chrono::NaiveDateTime;
use diesel::prelude::*;

#[derive(Debug, Serialize, Deserialize, Queryable)]
pub struct Todo {
    pub id: i32,
    pub title: String,
    pub completed: bool,
    pub order: i32,
    pub url: String,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

#[derive(Debug, Insertable, AsChangeset, Serialize, Deserialize)]
#[diesel(table_name = todos)]
pub struct TodoPayload {
    pub title: Option<String>,
    pub completed: Option<bool>,
    pub order: Option<i32>,
    pub url: Option<String>,
}

pub fn add(new_todo: TodoPayload, conn: &mut PgConnection) -> Result<Todo, DbError> {
    use crate::models::todo::todos::dsl::*;

    let res = diesel::insert_into(todos)
        .values(&new_todo)
        .get_result(conn);

    match res {
        Ok(todo) => Ok(todo),
        Err(e) => Err(Box::new(e)),
    }
}

pub fn all(conn: &mut PgConnection) -> Result<Vec<Todo>, DbError> {
    use crate::models::todo::todos::dsl::*;

    let res = todos.load::<Todo>(conn);

    match res {
        Ok(_) => Ok(res.unwrap()),
        Err(e) => Err(Box::new(e)),
    }
}

pub fn find(todo_id: i32, conn: &mut PgConnection) -> Result<Todo, DbError> {
    use crate::models::todo::todos::dsl::*;

    let res = todos
        .find(todo_id)
        .first::<Todo>(conn);

    match res {
        Ok(todo) => Ok(todo),
        Err(e) => Err(Box::new(e)),
    }
}

pub fn update(todo_id: i32, changes: TodoPayload, conn: &mut PgConnection) -> Result<Todo, DbError> {
    use crate::models::todo::todos::dsl::*;

    let res = diesel::update(todos.find(todo_id))
        .set(&changes)
        .get_result::<Todo>(conn);

    match res {
        Ok(todo) => Ok(todo),
        Err(e) => Err(Box::new(e)),
    }
}

pub fn delete(todo_id: i32, conn: &mut PgConnection) -> Result<Todo, DbError> {
    use crate::models::todo::todos::dsl::*;

    let res = diesel::delete(todos.find(todo_id))
        .get_result::<Todo>(conn);

    match res {
        Ok(todo) => Ok(todo),
        Err(e) => Err(Box::new(e)),
    }
}
