import boto3, os
from urllib2 import urlopen
from bs4 import BeautifulSoup

TOPIC = os.environ['TOPIC']

sns = boto3.client('sns')

def handler(event, context):
	html = urlopen('https://hub.docker.com/r/library/ubuntu/tags/').read()
	soup = BeautifulSoup(html, 'html.parser')
	if 'day' not in soup.find(string='latest').next_element.next_element.next_element.text:
		msg = 'Docker library/ubuntu: New version detected'
		print msg
		sns.publish(TopicArn=TOPIC, Subject=msg, Message=msg)
