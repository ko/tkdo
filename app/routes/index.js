import Ember from 'ember';

export default Ember.Route.extend({

    model: function() {
        return {
            data: this.store.findAll('task'),
            task: {}
        }
    },

    actions: {
        createTask(info) {

            console.log("kenko:createTask:1");
            let newTask = this.store.createRecord('task', {
                title: info.title,
                description: info.description,
                createdData: new Date(),
            });

            newTask.save();
        },

        deleteTask(info) {
            let task = this.store.find('task', info.id).then(function (task) {
                task.destroyRecord();  
            });;
        }
    }
});
