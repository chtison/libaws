import boto3, os
from urllib2 import urlopen
from bs4 import BeautifulSoup

TOPIC = os.environ['TOPIC']
GOVERSION = os.environ['GOVERSION']

sns = boto3.client('sns')

def handler(event, context):
	html = urlopen('https://golang.org/dl/').read()
	soup = BeautifulSoup(html, 'html.parser')
	version = soup.select('#stable + div')[0]['id']
	if version != GOVERSION:
		msg = 'Golang: New version detected: {}'.format(version)
		print msg
		sns.publish(TopicArn=TOPIC, Subject=msg, Message=msg)
