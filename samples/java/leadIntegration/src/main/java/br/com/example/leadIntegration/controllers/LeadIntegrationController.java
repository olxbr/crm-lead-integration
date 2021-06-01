package br.com.example.leadIntegration.controllers;

import org.springframework.web.bind.annotation.*;
import br.com.example.leadIntegration.models.Lead;

@RestController
@RequestMapping("/leads")
public class LeadIntegrationController {

    @GetMapping("/health-check")
    public String healthCheck() {
        return "{\"message\" : \"Ok\"}";
    }

    @PostMapping("/lead")
    public String lead(@RequestBody Lead lead){
        return "{\"message\": \"Lead successfully received\"}";
    }
}
