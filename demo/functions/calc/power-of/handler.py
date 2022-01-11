import json
import pprint


def response(output):
    _fal = {
        'success': output is not None,
        'output': str(output or '')
    }

    return { '_fal': json.dumps(_fal, separators=(',', ':')) }


def power_of(event, context):
    print('Got event: ' + pprint.pformat(event))

    if not event.get('_fal'):
        return response(None)

    try:
        fal = json.loads(event['_fal'])
        fal_input = int(fal['input'])

        return response(pow(fal_input, 2))

    except (TypeError, ValueError) as err:
        print('Could not perform operation: ' + str(err))

    return response(None)
