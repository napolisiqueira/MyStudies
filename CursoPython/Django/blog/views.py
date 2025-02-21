from django.http import HttpResponse

# Create your views here.
def blog(request):
    return HttpResponse('Uma mensagem ainda mais especial.')