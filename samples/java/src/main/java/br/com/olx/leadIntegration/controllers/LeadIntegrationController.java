package br.com.olx.leadIntegration.controllers;

import org.springframework.web.bind.annotation.*;
import br.com.olx.leadIntegration.models.Lead;

@RestController
@RequestMapping("/leads")
public class LeadIntegrationController {
    
    @PostMapping("/lead")
    public String lead(@RequestBody Lead lead){
        return "{\"message\": \"Lead successfully received\"}";
    }
}
