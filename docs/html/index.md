---
layout: default
permalink: /
title: LibAWS
rightbar:
  - example
title_append_brand: false
---

# What is LibAWS ?

Library Amazon Web Service is a template engine embedding predefined
templates for aws cloudformation.
{: .la-lead }

## Example
{: #example }

Specify parameters to the template engine via a [libaws.yaml](/reference/libaws)
file.
<div class="la-example">
<span class="la-example-name">libaws.yaml</span>
{% highlight yaml %}{% include examples/libaws.yaml %}{% endhighlight %}
</div>

Write template(s) where the [golang templating language](https://golang.org/pkg/text/template/)
will be executed.
<div class="la-example">
<span class="la-example-name">cloudformation.tmpl.yaml</span>
{% highlight yaml %}{% include examples/cloudformation.tmpl.yaml %}{% endhighlight %}
</div>

Write some data file(s) that you will pass to the template engine to fill
your template(s).
<div class="la-example">
<span class="la-example-name">cloudformation.data.yaml</span>
{% highlight yaml %}{% include examples/cloudformation.data.yaml %}{% endhighlight %}
</div>

You can then execute in the same directory as the 3 files:
<div class="la-example">
<span class="la-example-name">shell</span>
{% highlight sh %}$ libaws template run{% endhighlight %}
</div>

Which will output:
<div class="la-example">
<span class="la-example-name">stdout</span>
{% highlight yaml %}{% include examples/cloudformation.yaml %}{% endhighlight %}
</div>
