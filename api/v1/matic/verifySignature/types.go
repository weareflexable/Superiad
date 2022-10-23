package verifysignature

type VerifySignatureRequest struct {
	UserId    string `json:"userId" binding:"required"`
	Message   string `json:"message" binding:"required"`
	Signature string `json:"signature" binding:"required,hexadecimal,len=132"`
}

type VerifySignaturePayload struct {
	IsSignedByUser bool `json:"isSignedByUser" binding:"required"`
}
