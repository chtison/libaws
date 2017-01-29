import boto3, os
from urllib2 import urlopen
from bs4 import BeautifulSoup

TOPIC = os.environ['TOPIC']
VERSION = os.environ['VERSION']

sns = boto3.client('sns')

def handler(event, context):
	html = urlopen('https://golang.org/dl/').read()
	soup = BeautifulSoup(html, 'html.parser')
	version = soup.select('#stable + div')[0]['id']
	if version != VERSION:
		msg = 'Golang: New version detected: {}'.format(version)
		print msg
		print 'Publishing to SNS topic: {}'.format(TOPIC)
		sns.publish(TopicArn=TOPIC, Subject=msg, Message=msg)
