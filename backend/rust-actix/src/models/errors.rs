use actix_web::http::StatusCode;

pub type DbError = Box<dyn std::error::Error + Send + Sync>;

pub fn message(e: &DbError) -> String {
    match e.downcast_ref::<diesel::result::Error>() {
        Some(diesel::result::Error::NotFound) => "Record not found".to_string(),
        Some(diesel::result::Error::QueryBuilderError(_)) => "Invalid query".to_string(),
        _ => "Unknown database error".to_string(),
    }
}

pub fn status_code(e: &DbError) -> StatusCode {
    match e.downcast_ref::<diesel::result::Error>() {
        Some(diesel::result::Error::NotFound) => StatusCode::NOT_FOUND,
        Some(diesel::result::Error::QueryBuilderError(_)) => StatusCode::BAD_REQUEST,
        _ => StatusCode::INTERNAL_SERVER_ERROR,
    }
}
