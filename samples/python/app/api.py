from fastapi import FastAPI, Depends

from app.auth import verify_token
from app.schemas import LeadRequestSchema, ResponseSchema


app = FastAPI(
    title="Lead Integration - Python", description="A basic sample using python to show how to receive leads from zap+."
)


@app.get("/")
@app.get("/health-check")
def health_check():
    return {"message": "OK"}


@app.post("/lead", response_model=ResponseSchema)
def lead(lead: LeadRequestSchema, authorized: bool = Depends(verify_token)):
    return {"message": "Lead successfully received"}
