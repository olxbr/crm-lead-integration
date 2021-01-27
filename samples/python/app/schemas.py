from pydantic import BaseModel, Field


class LeadRequestSchema(BaseModel):
    lead_origin: str = Field(alias="leadOrigin")
    timestamp: str
    origin_lead_id: str = Field(alias="originLeadId")
    origin_listing_id: str = Field(alias="originListingId")
    client_listing_id: str = Field(alias="clientListingId")
    name: str
    email: str
    ddd: str
    phone: str
    message: str

    class Config:
        orm_mode = False
        allow_population_by_field_name = True


class ResponseSchema(BaseModel):
    message: str

    class Config:
        orm_mode = False
        allow_population_by_field_name = True
