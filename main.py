import click
import zmq
from time import time

@click.command
@click.option('--type', '-t', type=click.Choice(['sub'], case_sensitive=False), required=True, help='Type of ZeroMQ socket to use.')
@click.option('--address', '-a', type=str, help='Address to bind the socket to (protocol:ip:addr). e.g., tcp://localhost:5555')
def main(type: str, address: str):
    context = zmq.Context()
    if type == 'sub':
        zmq_type = zmq.SUB
        socket_options = {zmq.SUBSCRIBE: b''}
    else:
        click.echo(f"Unsupported socket type: {type}")
        return

    socket = context.socket(zmq_type)
    socket.connect(address)

    for option, value in socket_options.items():
        socket.setsockopt(option, value)

    if zmq_type == zmq.SUB:
        click.echo(f"Subscribed to {address} with socket type {type.upper()}")
        while True:
            try:
                message = socket.recv()
                click.echo(f"Receive POSIX time: {time()} | Message: {len(message)} bytes")
            except zmq.ZMQError as e:
                click.echo(f"ZMQ Error: {e}")
                break

if __name__ == "__main__":
    main()
