use actix_web::HttpResponse;
use actix_web::http::StatusCode;
use std::collections::HashMap;
use serde::Serialize;
use crate::models::errors;

pub fn json_response<T>(status: StatusCode, envelope: &str, body: T) -> HttpResponse
where
    T: Serialize,
{
    let body_response = HashMap::from([
        (envelope, body),
    ]);

    HttpResponse::build(status).json(body_response)
}

pub fn json_error_response(status: StatusCode, message: &str) -> HttpResponse {
    let body_response = HashMap::from([
        ("error", message),
    ]);

    HttpResponse::build(status).json(body_response)
}

pub fn json_error_db_response(e: errors::DbError) -> HttpResponse {
    let body_response = HashMap::from([
        ("error", errors::message(&e)),
    ]);

    let status = errors::status_code(&e);

    HttpResponse::build(status).json(body_response)
}
