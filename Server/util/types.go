package util

// AddToCartReq is the data to be added to a users cart
type AddToCartReq struct {
	UserID     string
	ItemList   []string
	ItemIDList []string
}

// StoreCheckoutReq is the /store/checkout req
type StoreCheckoutReq struct {
	UserID   string
	ItemList []string
}

// HomeCheckout has the data for /home/checkout
type HomeCheckout struct {
	UserID   string
	ItemList []string
}

// GetCartRes is the data in the response
type GetCartRes struct {
	ItemList   []string
	ItemIDList []string
}

// CartData is the data obtained from the DB cart
type CartData struct {
	UserID string
	Item   string
	ItemID string
}

type GetCartReq struct {
	UserID string
}

/*
{
  "responseId": "ea3d77e8-ae27-41a4-9e1d-174bd461b68c",
  "session": "projects/your-agents-project-id/agent/sessions/88d13aa8-2999-4f71-b233-39cbf3a824a0",
  "queryResult": {
    "queryText": "user's original query to your agent",
    "parameters": {
      "param": "param value"
    },
    "allRequiredParamsPresent": true,
    "fulfillmentText": "Text defined in Dialogflow's console for the intent that was matched",
    "fulfillmentMessages": [
      {
        "text": {
          "text": [
            "Text defined in Dialogflow's console for the intent that was matched"
          ]
        }
      }
    ],
    "outputContexts": [
      {
        "name": "projects/your-agents-project-id/agent/sessions/88d13aa8-2999-4f71-b233-39cbf3a824a0/contexts/generic",
        "lifespanCount": 5,
        "parameters": {
          "param": "param value"
        }
      }
    ],
    "intent": {
      "name": "projects/your-agents-project-id/agent/intents/29bcd7f8-f717-4261-a8fd-2d3e451b8af8",
      "displayName": "Matched Intent Name"
    },
    "intentDetectionConfidence": 1,
    "diagnosticInfo": {},
    "languageCode": "en"
  },
  "originalDetectIntentRequest": {}
}
*/

// GAPostReq is the POST sent by Dialogflow
type GAPostReq struct {
	ResponseId string
	session    string
}
