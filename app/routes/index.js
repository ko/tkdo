import Ember from 'ember';

export default Ember.Route.extend({

    model: function() {
        return {
            data: this.store.findAll('task'),
            task: {}
        };
    },

    actions: {
        createTask(info) {

            let newTask = this.store.createRecord('task', {
                title: info.title,
                description: info.description,
                createdDate: new Date(),
            });

            newTask.save();
        },

        deleteTask(info) {
            this.store.find('task', info.id).then(function (task) {
                task.destroyRecord();  
            });
        }
    }
});
