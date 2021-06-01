package br.com.example.leadIntegration.models;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
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