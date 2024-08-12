// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Spiderpool Agent",
    "title": "Spiderpool-Agent API",
    "version": "v1"
  },
  "basePath": "/v1",
  "paths": {
    "/coordinator/config": {
      "get": {
        "description": "Send a request to daemonset to get coordinator config\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Get coordinator config",
        "parameters": [
          {
            "name": "get-coordinator-config",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/GetCoordinatorArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/CoordinatorConfig"
            }
          },
          "500": {
            "description": "Config not initialized successfully or others",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      }
    },
    "/ipam/healthy": {
      "get": {
        "description": "Check spiderpool daemonset health to make sure whether it's ready\nfor CNI plugin usage\n",
        "tags": [
          "connectivity"
        ],
        "summary": "Get health of spiderpool daemon",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/ipam/ip": {
      "post": {
        "description": "Send a request to daemonset to ask for an ip assignment\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Get ip from spiderpool daemon",
        "parameters": [
          {
            "name": "ipam-add-args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/IpamAddArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/IpamAddResponse"
            }
          },
          "500": {
            "description": "Allocation failure",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      },
      "delete": {
        "description": "Send a request to daemonset to ask for an ip deleting\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Delete ip from spiderpool daemon",
        "parameters": [
          {
            "name": "ipam-del-args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/IpamDelArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Addresses release failure",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      }
    },
    "/ipam/ips": {
      "post": {
        "description": "Assign multiple ip for a pod, case for spiderflat compent\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Assign multiple ip as a batch",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Allocation failure",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      },
      "delete": {
        "description": "Delete multiple ip for a pod, case for spiderflat compent\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Delete multiple ip as a batch",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Addresses release failure",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      }
    },
    "/runtime/liveness": {
      "get": {
        "description": "Check pod liveness probe",
        "tags": [
          "runtime"
        ],
        "summary": "Liveness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/runtime/readiness": {
      "get": {
        "description": "Check pod readiness probe",
        "tags": [
          "runtime"
        ],
        "summary": "Readiness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/runtime/startup": {
      "get": {
        "description": "Check pod startup probe",
        "tags": [
          "runtime"
        ],
        "summary": "Startup probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/workloadendpoint": {
      "get": {
        "description": "Get workloadendpoint details for spiderflat use\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Get workloadendpoint status",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Get workloadendpoint failure"
          }
        }
      }
    }
  },
  "definitions": {
    "CoordinatorConfig": {
      "description": "Coordinator config",
      "type": "object",
      "required": [
        "overlayPodCIDR",
        "serviceCIDR",
        "mode",
        "tunePodRoutes"
      ],
      "properties": {
        "detectGateway": {
          "type": "boolean"
        },
        "detectIPConflict": {
          "type": "boolean"
        },
        "hijackCIDR": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "hostRPFilter": {
          "type": "integer"
        },
        "hostRuleTable": {
          "type": "integer"
        },
        "mode": {
          "type": "string"
        },
        "overlayPodCIDR": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "podDefaultRouteNIC": {
          "type": "string"
        },
        "podMACPrefix": {
          "type": "string"
        },
        "serviceCIDR": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "tunePodRoutes": {
          "type": "boolean"
        }
      }
    },
    "DNS": {
      "description": "IPAM CNI types DNS",
      "type": "object",
      "properties": {
        "domain": {
          "type": "string"
        },
        "nameservers": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "options": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "search": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Error": {
      "description": "API error",
      "type": "string"
    },
    "GetCoordinatorArgs": {
      "description": "Get Coordinator Args",
      "type": "object",
      "properties": {
        "podName": {
          "type": "string"
        },
        "podNamespace": {
          "type": "string"
        }
      }
    },
    "IpConfig": {
      "description": "IPAM IPs struct, contains ifName, Address and Gateway",
      "type": "object",
      "required": [
        "version",
        "address",
        "nic"
      ],
      "properties": {
        "address": {
          "type": "string"
        },
        "gateway": {
          "type": "string"
        },
        "ipPool": {
          "type": "string"
        },
        "nic": {
          "type": "string"
        },
        "version": {
          "type": "integer",
          "enum": [
            4,
            6
          ]
        },
        "vlan": {
          "type": "integer"
        }
      }
    },
    "IpamAddArgs": {
      "description": "IPAM request args",
      "type": "object",
      "required": [
        "containerID",
        "ifName",
        "netNamespace",
        "podNamespace",
        "podName",
        "podUID"
      ],
      "properties": {
        "cleanGateway": {
          "type": "boolean"
        },
        "containerID": {
          "type": "string"
        },
        "defaultIPv4IPPool": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "defaultIPv6IPPool": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ifName": {
          "type": "string"
        },
        "netNamespace": {
          "type": "string"
        },
        "podName": {
          "type": "string"
        },
        "podNamespace": {
          "type": "string"
        },
        "podUID": {
          "type": "string"
        }
      }
    },
    "IpamAddResponse": {
      "description": "IPAM assignment IPs information",
      "type": "object",
      "required": [
        "ips"
      ],
      "properties": {
        "dns": {
          "type": "object",
          "$ref": "#/definitions/DNS"
        },
        "ips": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/IpConfig"
          }
        },
        "routes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Route"
          }
        }
      }
    },
    "IpamDelArgs": {
      "description": "IPAM release IP information",
      "type": "object",
      "required": [
        "containerID",
        "ifName",
        "podNamespace",
        "podName",
        "podUID"
      ],
      "properties": {
        "containerID": {
          "type": "string"
        },
        "ifName": {
          "type": "string"
        },
        "netNamespace": {
          "type": "string"
        },
        "podName": {
          "type": "string"
        },
        "podNamespace": {
          "type": "string"
        },
        "podUID": {
          "type": "string"
        }
      }
    },
    "Route": {
      "description": "IPAM CNI types Route",
      "type": "object",
      "required": [
        "ifName",
        "dst",
        "gw"
      ],
      "properties": {
        "dst": {
          "type": "string"
        },
        "gw": {
          "type": "string"
        },
        "ifName": {
          "type": "string"
        }
      }
    }
  },
  "x-schemes": [
    "unix"
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Spiderpool Agent",
    "title": "Spiderpool-Agent API",
    "version": "v1"
  },
  "basePath": "/v1",
  "paths": {
    "/coordinator/config": {
      "get": {
        "description": "Send a request to daemonset to get coordinator config\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Get coordinator config",
        "parameters": [
          {
            "name": "get-coordinator-config",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/GetCoordinatorArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/CoordinatorConfig"
            }
          },
          "500": {
            "description": "Config not initialized successfully or others",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      }
    },
    "/ipam/healthy": {
      "get": {
        "description": "Check spiderpool daemonset health to make sure whether it's ready\nfor CNI plugin usage\n",
        "tags": [
          "connectivity"
        ],
        "summary": "Get health of spiderpool daemon",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/ipam/ip": {
      "post": {
        "description": "Send a request to daemonset to ask for an ip assignment\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Get ip from spiderpool daemon",
        "parameters": [
          {
            "name": "ipam-add-args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/IpamAddArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/IpamAddResponse"
            }
          },
          "500": {
            "description": "Allocation failure",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      },
      "delete": {
        "description": "Send a request to daemonset to ask for an ip deleting\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Delete ip from spiderpool daemon",
        "parameters": [
          {
            "name": "ipam-del-args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/IpamDelArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Addresses release failure",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      }
    },
    "/ipam/ips": {
      "post": {
        "description": "Assign multiple ip for a pod, case for spiderflat compent\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Assign multiple ip as a batch",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Allocation failure",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      },
      "delete": {
        "description": "Delete multiple ip for a pod, case for spiderflat compent\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Delete multiple ip as a batch",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Addresses release failure",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      }
    },
    "/runtime/liveness": {
      "get": {
        "description": "Check pod liveness probe",
        "tags": [
          "runtime"
        ],
        "summary": "Liveness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/runtime/readiness": {
      "get": {
        "description": "Check pod readiness probe",
        "tags": [
          "runtime"
        ],
        "summary": "Readiness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/runtime/startup": {
      "get": {
        "description": "Check pod startup probe",
        "tags": [
          "runtime"
        ],
        "summary": "Startup probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/workloadendpoint": {
      "get": {
        "description": "Get workloadendpoint details for spiderflat use\n",
        "tags": [
          "daemonset"
        ],
        "summary": "Get workloadendpoint status",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Get workloadendpoint failure"
          }
        }
      }
    }
  },
  "definitions": {
    "CoordinatorConfig": {
      "description": "Coordinator config",
      "type": "object",
      "required": [
        "overlayPodCIDR",
        "serviceCIDR",
        "mode",
        "tunePodRoutes"
      ],
      "properties": {
        "detectGateway": {
          "type": "boolean"
        },
        "detectIPConflict": {
          "type": "boolean"
        },
        "hijackCIDR": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "hostRPFilter": {
          "type": "integer"
        },
        "hostRuleTable": {
          "type": "integer"
        },
        "mode": {
          "type": "string"
        },
        "overlayPodCIDR": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "podDefaultRouteNIC": {
          "type": "string"
        },
        "podMACPrefix": {
          "type": "string"
        },
        "serviceCIDR": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "tunePodRoutes": {
          "type": "boolean"
        }
      }
    },
    "DNS": {
      "description": "IPAM CNI types DNS",
      "type": "object",
      "properties": {
        "domain": {
          "type": "string"
        },
        "nameservers": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "options": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "search": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Error": {
      "description": "API error",
      "type": "string"
    },
    "GetCoordinatorArgs": {
      "description": "Get Coordinator Args",
      "type": "object",
      "properties": {
        "podName": {
          "type": "string"
        },
        "podNamespace": {
          "type": "string"
        }
      }
    },
    "IpConfig": {
      "description": "IPAM IPs struct, contains ifName, Address and Gateway",
      "type": "object",
      "required": [
        "version",
        "address",
        "nic"
      ],
      "properties": {
        "address": {
          "type": "string"
        },
        "gateway": {
          "type": "string"
        },
        "ipPool": {
          "type": "string"
        },
        "nic": {
          "type": "string"
        },
        "version": {
          "type": "integer",
          "enum": [
            4,
            6
          ]
        },
        "vlan": {
          "type": "integer"
        }
      }
    },
    "IpamAddArgs": {
      "description": "IPAM request args",
      "type": "object",
      "required": [
        "containerID",
        "ifName",
        "netNamespace",
        "podNamespace",
        "podName",
        "podUID"
      ],
      "properties": {
        "cleanGateway": {
          "type": "boolean"
        },
        "containerID": {
          "type": "string"
        },
        "defaultIPv4IPPool": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "defaultIPv6IPPool": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ifName": {
          "type": "string"
        },
        "netNamespace": {
          "type": "string"
        },
        "podName": {
          "type": "string"
        },
        "podNamespace": {
          "type": "string"
        },
        "podUID": {
          "type": "string"
        }
      }
    },
    "IpamAddResponse": {
      "description": "IPAM assignment IPs information",
      "type": "object",
      "required": [
        "ips"
      ],
      "properties": {
        "dns": {
          "type": "object",
          "$ref": "#/definitions/DNS"
        },
        "ips": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/IpConfig"
          }
        },
        "routes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Route"
          }
        }
      }
    },
    "IpamDelArgs": {
      "description": "IPAM release IP information",
      "type": "object",
      "required": [
        "containerID",
        "ifName",
        "podNamespace",
        "podName",
        "podUID"
      ],
      "properties": {
        "containerID": {
          "type": "string"
        },
        "ifName": {
          "type": "string"
        },
        "netNamespace": {
          "type": "string"
        },
        "podName": {
          "type": "string"
        },
        "podNamespace": {
          "type": "string"
        },
        "podUID": {
          "type": "string"
        }
      }
    },
    "Route": {
      "description": "IPAM CNI types Route",
      "type": "object",
      "required": [
        "ifName",
        "dst",
        "gw"
      ],
      "properties": {
        "dst": {
          "type": "string"
        },
        "gw": {
          "type": "string"
        },
        "ifName": {
          "type": "string"
        }
      }
    }
  },
  "x-schemes": [
    "unix"
  ]
}`))
}
