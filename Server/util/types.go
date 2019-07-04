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

// CartItem is a single cart item
type CartItem struct {
	Item   string
	ItemID string
}

// // GetCartRes is the data in the response
// type GetCartRes struct {
// 	ItemList   []string
// 	ItemIDList []string
// }

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

// Recipe is a single recipe
type Recipe struct {
	RecipeName   string
	MissingItems []CartItem
}

// MLGetReq is the ML servers POST req
type MLGetReq struct {
	Recipes []Recipe
}

// PartRecipe is the recipe based on the previous items
type PartRecipe struct {
	RecipeName string
	ItemsList  []string
}

// MLPostRes responds to client
type MLPostRes struct {
	AddedCartList     []CartItem
	PredictedCartList []CartItem
	// Items             []PartRecipe
	Recipes []Recipe
}

/*
{
    "id": "0263991f-b4f8-47a4-84b4-e011bcb2a905",
    "project": "accf9d3c-0d15-4967-97fa-428640c2cf37",
    "iteration": "2d1880f3-886d-49a6-a03f-caad21ae36d7",
    "created": "2019-07-03T04:54:34.827Z",
    "predictions": [
        {
            "probability": 0.9988686,
            "tagId": "8d18fac3-4f2e-41c7-a87b-ecb6b4dca2aa",
            "tagName": "Pizza"
        },
        {
            "probability": 0.000939635152,
            "tagId": "acea3116-aac9-4c26-ae9f-9a7e44320e66",
            "tagName": "nachos"
        },
        {
            "probability": 0.000177755355,
            "tagId": "30120d8c-5112-449e-9a0f-2b67ba131b33",
            "tagName": "sphagetti"
        },
        {
            "probability": 0.00000711493476,
            "tagId": "a74943b1-1c71-410f-a697-4437683a0045",
            "tagName": "tacos"
        },
        {
            "probability": 0.00000691846753,
            "tagId": "c8871228-197d-4f3e-93a0-7962d41bf024",
            "tagName": "pancakes"
        }
    ]
}
*/

// Prediction is a single response
type Prediction struct {
	Probability float64
	TagID       string
	TagName     string
}

// ImageRecognitionRes response from image recognition server
type ImageRecognitionRes struct {
	ID          string
	Project     string
	Iteration   string
	Created     string
	Predictions []Prediction
}

// ImageRecCartRes is the hasher
type ImageRecCartRes struct {
	HashItems map[string][]string
}
