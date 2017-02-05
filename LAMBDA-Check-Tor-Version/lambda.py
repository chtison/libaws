import boto3, os
from urllib2 import urlopen

TOPIC = os.environ['TOPIC']
VERSION = os.environ['VERSION']

sns = boto3.client('sns')

def handler(event, context):
	html = urlopen('https://gitweb.torproject.org/tor.git/plain/ReleaseNotes').read()
	i1 = html.index('Changes in version')
	i2 = html.index('\n', i1)
	version = html[i1:i2].split()[3]
	if version != VERSION:
		msg = 'Tor: New version detected: {}'.format(version)
		print msg
		sns.publish(TopicArn=TOPIC, Subject=msg, Message=msg)
