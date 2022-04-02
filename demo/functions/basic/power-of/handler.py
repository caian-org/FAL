import json


def power_of(event, context):
    print(f'Got event:\n {json.dumps(event, indent=2)}')

    fal_key = '_fal'
    if fal_key not in event:
        raise KeyError(f'Event is missing "{fal_key}" property')

    fal_input = int(event.get(fal_key))
    fal_output = pow(fal_input, 2)

    return {
        fal_key: str(fal_output)
    }
