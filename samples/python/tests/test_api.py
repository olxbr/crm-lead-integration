from json import dumps
from unittest import TestCase
from _pytest.monkeypatch import MonkeyPatch

from fastapi.testclient import TestClient

from app.api import app
from app.schemas import LeadRequestSchema

class ApiTest(TestCase):
    def setUp(self):
        self.monkeypatch = MonkeyPatch()
        self.client_api = TestClient(app)
        self.lead_request_model = LeadRequestSchema(lead_origin="VivaReal",
                                                    timestamp="2017-10-23T15:50:30.619Z",
                                                    origin_lead_id="59ee0fc6e4b043e1b2a6d863",
                                                    origin_listing_id="87027856",
                                                    client_listing_id="a40171",
                                                    name="Nome Consumidor",
                                                    email="nome.consumidor@email.com",
                                                    ddd="11",
                                                    phone="999999999",
                                                    message="lá, tenho interesse neste imóvel. Aguardo o contato. Obrigado.")

    def test_when_I_call_health_check_should_be_success(self):
        response = self.client_api.get("/health-check")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json(), {"message": "OK"})

    def test_when_I_call_health_check_using_a_root_path_should_be_success(self):
        response = self.client_api.get("/")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json(), {"message": "OK"})

    def test_when_I_call_lead_using_a_wrong_credential_should_be_error(self):
        response = self.client_api.post("/lead", json=self.lead_request_model.json(), headers={"Authorization": "bla"})
        self.assertEqual(response.status_code, 401)

    def test_when_I_call_lead_using_a_right_credential_should_be_success(self):
        self.monkeypatch.setenv("SECRET_KEY", "594F803B380A41396ED63DCA39503542")
        response = self.client_api.post("/lead", json=self.lead_request_model.dict(), headers={"Authorization": "dml2YXJlYWw6NTk0RjgwM0IzODBBNDEzOTZFRDYzRENBMzk1MDM1NDI="})
        self.assertEqual(response.status_code, 200)
