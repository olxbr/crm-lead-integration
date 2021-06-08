package br.com.example.leadIntegration;

import br.com.example.leadIntegration.models.Lead;
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

		final String baseUrl = "http://localhost:"+port+"/leads/health-check";
		URI uri = new URI(baseUrl);

		ResponseEntity<String> result = this.restTemplate.getForEntity(uri, String.class);

		Assert.assertEquals(200, result.getStatusCodeValue());
		Assert.assertEquals("{\"message\": \"OK\"}", result.getBody());

	}

	@Test
	public void recieveLeadShouldBeSuccess() throws Exception {

		final String baseUrl = "http://localhost:"+port+"/leads/lead";
		URI uri = new URI(baseUrl);

		Lead lead = new Lead("VivaReal", "2017-10-23T15:50:30.619Z", "59ee0fc6e4b043e1b2a6d863", "87027856",
				"a40171", "Nome Consumidor", "nome.consumidor@email.com", "11", "999999999",
				"Olá, tenho interesse neste imóvel. Aguardo o contato. Obrigado.");

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
	public void recieveLeadShouldFailWihtoutAuth() throws Exception {

		final String baseUrl = "http://localhost:"+port+"/leads/lead";
		URI uri = new URI(baseUrl);

		Lead lead = new Lead("VivaReal", "2017-10-23T15:50:30.619Z", "59ee0fc6e4b043e1b2a6d863", "87027856",
				"a40171", "Nome Consumidor", "nome.consumidor@email.com", "11", "999999999",
				"Olá, tenho interesse neste imóvel. Aguardo o contato. Obrigado.");

		HttpHeaders headers = new HttpHeaders();
		headers.set("Authorization", "");

		HttpEntity<Lead> request = new HttpEntity<>(lead, headers);

		ResponseEntity<String> result = this.restTemplate.postForEntity(uri, request, String.class);

		Assert.assertEquals(401, result.getStatusCodeValue());
	}
}
