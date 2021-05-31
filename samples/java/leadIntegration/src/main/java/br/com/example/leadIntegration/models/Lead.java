package br.com.example.leadIntegration.models;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class Lead {
    private String leadOrigin;
    private String timestamp;
    private String originLeadId;
    private String originListingId;
    private String clientListingId;
    private String name;
    private String email;
    private String ddd;
    private String phone;
    private String message;
}