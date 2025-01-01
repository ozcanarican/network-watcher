#!/bin/bash
echo "Switching to powersafe mode"
nohup qm start 113 &
nohup qm start 101 &
nohup pct start 106 &
nohup pct start 120 &
nohup pct start 121 &
nohup pct start 122 &
nohup pct start 123 &
nohup pct start 124 &
nohup pct start 132 &
