# Scores Project

## Getting Started

This project uses Docker and Kubernetes to manage and deploy the `scores` application. Follow the instructions below to get the application up and running.

### Prerequisites

- Docker installed on your machine
- Kubernetes installed and configured (e.g., Minikube or Docker Desktop with Kubernetes enabled)
- A Docker Hub account

### Running the Project

#### Step 1: Build and Push Docker Image

1. **Login to Docker Hub:**
   ```sh
   docker login
Enter your Docker Hub username and password.

2. **Build the Docker image:**
```sh
docker build -t scores:latest .
```

3. **Deploy to Kubernetes:**
```shell
cd /path/to/project/go-scores/k8s
```

4. **Apply the Kubernetes manifests:**
```shell
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f ingress.yaml
```

## gRPC API

### Aggregated category scores over a period of time
E.g. what have the daily ticket scores been for a past week or what were the scores between 1st and 31st of January.

For periods longer than one month weekly aggregates should be returned instead of daily values.

#### Endpoint:
* **host**: scores.local 
* **method**: POST 
* **pathAndQuery**: /get_by_categories.GetByCategories/GetAggregatedScores 
* **port**: 80 
* **scheme**: http 

#### Request Message:
```json
{
  "from": "2019-08-01", //Y-m-d
  "to": "2019-08-07"    //Y-m-d
}
```

#### Response Message:

```json
{
  "meta": {
    "periods": [
      {
        "id": 1,
        "start": 1564617600,
        "end": 1564617600
      }
    ],
    "categories": [
      {
        "id": 1,
        "name": "GDPR"
      },
      {
        "id": 2,
        "name": "Spelling"
      }
    ]
  },
  "data": [
    {
      "category_id": 1,
      "num_of_reviews": 120,
      "periods": [
        {
          "id": 1,
          "score": 40.5
        }
      ],
      "total_score": 40.5
    },
    {
      "category_id": 2,
      "num_of_reviews": 75,
      "periods": [
        {
          "id": 1,
          "score": 32.8
        }
      ],
      "total_score": 32.8
    }
  ]
}
```

### Scores by ticket
Aggregate scores for categories within defined period by ticket.

E.g. what aggregate category scores tickets have within defined rating time range have.

#### Endpoint:
* **host**: scores.local 
* **method**: POST 
* **pathAndQuery**: /get_by_tickets.GetByTickets/GetAggregatedScoresByTicket 
* **port**: 80 
* **scheme**: http 

Request Message:
```json
{
  "from": "2019-08-01", //Y-m-d
  "to": "2019-08-07"    //Y-m-d
}
```

#### Response Message:
```json
{
  "meta": {
    "categories": [
      {
        "id": 1,
        "name": "GDPR"
      },
      {
        "id": 2,
        "name": "Spelling"
      }
    ]
  },
  "data": [
    {
      "ticket_id": 1001,
      "categories": [
        {
          "id": 1,
          "score": 43.7
        },
        {
          "id": 2,
          "score": 34.9
        }
      ]
    },
    {
      "ticket_id": 1002,
      "categories": [
        {
          "id": 1,
          "score": 44.2
        },
        {
          "id": 2,
          "score": 42.0
        }
      ]
    }
  ]
}

```
### Period over Period score change

What has been the change from selected period over previous period.

E.g. current week vs. previous week or December vs. January change in percentages.

#### Endpoint:
* **host**: scores.local 
* **method**: POST 
* **pathAndQuery**: /get_changes_by_periods.GetChangesByPeriods/GetChangesByPeriods 
* **port**: 80 
* **scheme**: http 

Request Message:
```json
{
  "period": "week", //Enum ["week", "month", "year"]
}
```

#### Response Message:
```json
{
   "meta": {
      "periods": [
         {
            "id": 1,
            "start": 1564617600,
            "end": 1564617600
         },
         {
            "id": 2,
            "start": 1564790400,
            "end": 1564790400
         }
      ]
   },
   "data": [
      {
         "period_id": 2,
         "score_delta": 100.0
      },
      {
         "period_id": 1,
         "score_delta": -8.2
      }
   ]
}
````

### Overall quality score

What is the overall aggregate score for a period.

E.g. the overall score over past week has been 96%.

#### Endpoint:
* **host**: scores.local 
* **method**: POST 
* **pathAndQuery**: /get_overall.GetOverall/GetOverallScore 
* **port**: 80 
* **scheme**: http 

Request Message:
```json
{
  "from": "2019-08-01", //Y-m-d
  "to": "2019-08-07"    //Y-m-d
}
```

#### Response Message:
```json
{
  "field": 41.5
}
```