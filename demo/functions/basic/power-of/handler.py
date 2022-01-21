import json
import pprint


def response(output):
    _fal = {
        'success': output is not None,
        'output': str(output or '')
    }

    return {'_fal': json.dumps(_fal, separators=(',', ':'))}


def power_of(event, context):
    print('Got event: ' + pprint.pformat(event))
    fal_key = '_fal'

    try:
        if not event.get(fal_key):
            raise KeyError(f'Event is missing "{fal_key}" property')

        fal = json.loads(event[fal_key])
        fal_input = int(fal['input'])

        return response(pow(fal_input, 2))

    except Exception as err:
        print('Could not perform operation: ' + str(err))

    return response(None)
