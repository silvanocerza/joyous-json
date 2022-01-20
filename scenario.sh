set -e

cat yesterday.json_dump \
    | ./jj filter in team team-x \
    | ./jj add incident_id 6502