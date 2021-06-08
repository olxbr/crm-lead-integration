package br.com.olx.leadIntegration.models;

import lombok.*;

@Builder
@Data
public class Lead {
    @NonNull
    private String leadOrigin;
    @NonNull
    private String timestamp;
    @NonNull
    private String originLeadId;
    @NonNull
    private String originListingId;
    @NonNull
    private String clientListingId;
    @NonNull
    private String name;
    @NonNull
    private String email;
    @NonNull
    private String ddd;
    @NonNull
    private String phone;
    @NonNull
    private String message;
}