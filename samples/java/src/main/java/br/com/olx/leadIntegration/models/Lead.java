package br.com.olx.leadIntegration.models;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
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