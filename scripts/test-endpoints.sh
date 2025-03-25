#!/bin/bash

# Couleurs et styles
BOLD='\033[1m'
BLUE='\033[34m'
GREEN='\033[32m'
YELLOW='\033[33m'
RED='\033[31m'
RESET='\033[0m'

# Header
echo -e "${BOLD}${BLUE}┌──────────────────────────────────────────┐${RESET}"
echo -e "${BOLD}${BLUE}│           GOMON ENDPOINTS TEST           │${RESET}"
echo -e "${BOLD}${BLUE}└──────────────────────────────────────────┘${RESET}\n"

# Health Endpoint
echo -e "${BOLD}${GREEN}▶ Testing Health Endpoint${RESET}"
echo -e "${YELLOW}GET /health${RESET}"
response=$(curl -s -i http://localhost:8080/health)
echo "$response" | sed 's/^/  /'
echo -e "\n"

# Stats Endpoint
echo -e "${BOLD}${GREEN}▶ Testing Stats Endpoint${RESET}"
echo -e "${YELLOW}GET /stats${RESET}"
response=$(curl -s -i http://localhost:8080/stats)
echo "$response" | sed 's/^/  /' | python3 -m json.tool 2>/dev/null || echo "$response" | sed 's/^/  /'
echo -e "\n"

# Prometheus Metrics
echo -e "${BOLD}${GREEN}▶ Testing Prometheus Metrics${RESET}"
echo -e "${YELLOW}GET /metrics${RESET}"
response=$(curl -s -i http://localhost:8080/metrics)
echo "$response" | sed 's/^/  /'
echo -e "\n"

# Footer
echo -e "${BOLD}${BLUE}┌──────────────────────────────────────────┐${RESET}"
echo -e "${BOLD}${BLUE}│             TEST COMPLETE               │${RESET}"
echo -e "${BOLD}${BLUE}└──────────────────────────────────────────┘${RESET}" 