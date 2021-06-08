package br.com.olx.leadIntegration;

import br.com.olx.leadIntegration.models.Lead;
import org.junit.Assert;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.context.SpringBootTest.WebEnvironment;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseEntity;

import java.net.URI;
import java.util.Base64;

@SpringBootTest(webEnvironment = WebEnvironment.RANDOM_PORT)
class LeadIntegrationApplicationTests {

	@LocalServerPort
	private int port;

	@Autowired
	private TestRestTemplate restTemplate;

	@Value("${spring.security.user.name}")
	private String userName;

	@Value("${spring.security.user.password}")
	private String userPassword;

	@Test
	public void healthCheckShouldReturnOk() throws Exception {

		final String baseUrl = "http://localhost:"+port+"/actuator/health";
		URI uri = new URI(baseUrl);

		ResponseEntity<String> result = this.restTemplate.getForEntity(uri, String.class);

		Assert.assertEquals(200, result.getStatusCodeValue());
		Assert.assertEquals("{\"status\":\"UP\"}", result.getBody());

	}

	@Test
	public void recieveLeadShouldBeSuccess() throws Exception {

		final String baseUrl = "http://localhost:"+port+"/leads/lead";
		URI uri = new URI(baseUrl);

		Lead lead = Lead.builder()
				.leadOrigin("VivaReal")
				.timestamp("2017-10-23T15:50:30.619Z")
				.originLeadId("59ee0fc6e4b043e1b2a6d863")
				.originListingId("87027856")
				.clientListingId("a40171")
				.name("Nome Consumidor")
				.email("nome.consumidor@email.com")
				.ddd("11")
				.phone("999999999")
				.message("Olá, tenho interesse neste imóvel. Aguardo o contato. Obrigado.")
				.build();

		String userpass =  userName + ":" + userPassword;
		String auth = "Basic " + Base64.getEncoder().encodeToString(userpass.getBytes());

		HttpHeaders headers = new HttpHeaders();
		headers.set("Authorization", auth);

		HttpEntity<Lead> request = new HttpEntity<>(lead, headers);

		ResponseEntity<String> result = this.restTemplate.postForEntity(uri, request, String.class);

		Assert.assertEquals(200, result.getStatusCodeValue());
		Assert.assertEquals("{\"message\": \"Lead successfully received\"}", result.getBody());

	}

	@Test
	public void recieveLeadShouldFailWihtWrongAuth() throws Exception {

		final String baseUrl = "http://localhost:"+port+"/leads/lead";
		URI uri = new URI(baseUrl);

		Lead lead = Lead.builder()
				.leadOrigin("VivaReal")
				.timestamp("2017-10-23T15:50:30.619Z")
				.originLeadId("59ee0fc6e4b043e1b2a6d863")
				.originListingId("87027856")
				.clientListingId("a40171")
				.name("Nome Consumidor")
				.email("nome.consumidor@email.com")
				.ddd("11")
				.phone("999999999")
				.message("Olá, tenho interesse neste imóvel. Aguardo o contato. Obrigado.")
				.build();

		String userpass = "vivareal:A1b2C3d4";
		String auth = "Basic " + Base64.getEncoder().encodeToString(userpass.getBytes());

		HttpHeaders headers = new HttpHeaders();
		headers.set("Authorization", auth);

		HttpEntity<Lead> request = new HttpEntity<>(lead, headers);

		ResponseEntity<String> result = this.restTemplate.postForEntity(uri, request, String.class);

		Assert.assertEquals(401, result.getStatusCodeValue());
	}

	@Test
	public void recieveLeadShouldFailWihtoutAuth() throws Exception {

		final String baseUrl = "http://localhost:"+port+"/leads/lead";
		URI uri = new URI(baseUrl);

		Lead lead = Lead.builder()
				.leadOrigin("VivaReal")
				.timestamp("2017-10-23T15:50:30.619Z")
				.originLeadId("59ee0fc6e4b043e1b2a6d863")
				.originListingId("87027856")
				.clientListingId("a40171")
				.name("Nome Consumidor")
				.email("nome.consumidor@email.com")
				.ddd("11")
				.phone("999999999")
				.message("Olá, tenho interesse neste imóvel. Aguardo o contato. Obrigado.")
				.build();

		HttpHeaders headers = new HttpHeaders();
		headers.set("Authorization", "");

		HttpEntity<Lead> request = new HttpEntity<>(lead, headers);

		ResponseEntity<String> result = this.restTemplate.postForEntity(uri, request, String.class);

		Assert.assertEquals(401, result.getStatusCodeValue());
	}
}
