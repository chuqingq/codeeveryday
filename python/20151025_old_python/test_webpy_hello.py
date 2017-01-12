import web

render = web.template.render('templates/')

urls = ('/', 'Index')

class Index():
    def GET(self):
        ## return 'hello world'
        name = '<Bob>'
        return render.index(name)

if __name__ == '__main__':
    ## print globals()
    
    name = 'Bob'    
    app = web.application(urls, globals()) 
    app.run()
