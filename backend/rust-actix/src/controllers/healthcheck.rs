use actix_web::{get, HttpResponse, Responder};
use serde::Serialize;

#[derive(Serialize)]
pub struct HealtcheckResponse {
    pub status: String,
    pub message: String,
}

#[get("/healthcheck")]
async fn index() -> impl Responder {
    const MESSAGE: &str = "Simple TODO API with Rust and Actix Web";

    let response_json = &HealtcheckResponse {
        status: "available".to_string(),
        message: MESSAGE.to_string(),
    };

    HttpResponse::Ok().json(response_json)
}
