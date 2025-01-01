#!/bin/bash
echo "Switching to powersafe mode"
nohup qm shutdown 113 &
nohup qm shutdown 101 &
pct shutdown 106
pct shutdown 120
pct shutdown 121
pct shutdown 122
pct shutdown 123
pct shutdown 124
pct shutdown 132
